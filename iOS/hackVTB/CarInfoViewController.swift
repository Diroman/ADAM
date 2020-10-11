//
//  CarInfoViewController.swift
//  hackVTB
//
//  Created by andarbek on 11.10.2020.
//

import UIKit
import SwiftyJSON
import AACarousel
import Kingfisher
import Alamofire

class CarInfoViewController: UIViewController, AACarouselDelegate {
    var favoriteInt: Int?
    var jsonData: JSON?
    var urlList = [String]()
    var imageList = [UIImage]()
    var titleArray = [String]()
    
    
    @IBOutlet weak var titleCar: UILabel!
    
    @IBOutlet weak var priceCar: UILabel!
    @IBOutlet weak var modelCar: UILabel!
    
    @IBOutlet weak var countryCar: UILabel!

    
    @IBOutlet weak var like: UIButton!
    @IBOutlet weak var carouselView: AACarousel!
    @IBOutlet weak var info: UILabel!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        like.tintColor = .gray
        
        print(jsonData)
        getUrls()
        createCarousel()
        insertData()
    }
    
    @IBAction func addInFav(_ sender: UIButton) {
        sendLike()
    }
    
    @IBAction func calculator(_ sender: UIButton) {
        self.performSegue(withIdentifier: "showCacl", sender: Any?.self)
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?)
    {
        if segue.identifier == "showCacl" {
            if let nextVC = segue.destination as? CalculatorViewController {
                nextVC.price = self.priceCar.text
            }
        }
    }
    
    func sendLike(){
        if self.like.tintColor == .gray {
            let parameters: [String: Any] = [
                "UserID": 1,
                "JsonText" : jsonData?.description
            ]
            
            let headers: HTTPHeaders = [
                "x-access-token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIk5hbWUiOiIiLCJFbWFpbCI6InJvbWFhY3RhcG92QGdtYWlsLmNvbSIsImV4cCI6MTYwODM5Njg3Mn0.ZcQXzPSXv17hssstgw_oH-hsyzlidVx09kPQopig5BA"
                ]
            
            AF.request("http://34.123.226.60:8080/auth/favorite", method: .post, parameters: parameters as Parameters, encoding: JSONEncoding.default, headers: headers).responseJSON(completionHandler: { [weak self] response in
                do{
                let json = try JSON(data: response.data!)
                    self!.favoriteInt = json["Value"]["ID"].int
                    DispatchQueue.main.async {
                        self!.like.tintColor = .red
                    }
                } catch{
                    print("df")
                }
                
            }).resume()
        } else {
            let headers: HTTPHeaders = [
                "x-access-token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIk5hbWUiOiIiLCJFbWFpbCI6InJvbWFhY3RhcG92QGdtYWlsLmNvbSIsImV4cCI6MTYwODM5Njg3Mn0.ZcQXzPSXv17hssstgw_oH-hsyzlidVx09kPQopig5BA"
                ]
            
            AF.request("http://34.123.226.60:8080/auth/favorite/\(String(describing: self.favoriteInt!))", method: .delete, encoding: JSONEncoding.default, headers: headers).responseJSON(completionHandler: { [weak self] response in
                print(response)
                DispatchQueue.main.async {
                    self!.like.tintColor = .gray
                }
            }).resume()
        }

    }
    
    func createCarousel(){
        let pathArray = urlList
        carouselView.delegate = self
        carouselView.setCarouselData(paths: pathArray,  describedTitle: titleArray, isAutoScroll: true, timer: 2.0, defaultImage: "defaultImage")
        //optional methods
        carouselView.setCarouselOpaque(layer: false, describedTitle: false, pageIndicator: false)
        carouselView.setCarouselLayout(displayStyle: 2, pageIndicatorPositon: 5, pageIndicatorColor: nil, describedTitleColor: nil, layerColor: nil)
    }
    
    
    func getUrls(){
        for i in jsonData!["Cars"][0]["renderPhotos"]{
            if let url = i.1["path"].string{
                if url != "" {
                    self.urlList.append(url)
                    self.titleArray.append(jsonData!["Cars"][0]["Title"].string!)
                }
            }
        }
    }
    
    func insertData(){
        self.titleCar.text = jsonData!["Cars"][0]["Title"].string
        self.priceCar.text = jsonData!["Cars"][0]["prettyPrice"].string
        self.modelCar.text = jsonData!["Cars"][0]["Title"].string
        self.countryCar.text = jsonData!["Cars"][0]["country"].string
        self.info.text = jsonData!["Cars"][0]["info"].string
    }
    
    func downloadImages(_ url: String, _ index: Int) {
        
        let imageView = UIImageView()
        imageView.kf.setImage(with: URL(string: url)!, placeholder: UIImage.init(named: "defaultImage"), options: [.transition(.fade(0))], progressBlock: nil, completionHandler: { (downloadImage, error, cacheType, url) in
            self.carouselView.images[index] = downloadImage!
        })
    }
    
    func didSelectCarouselView(_ view: AACarousel ,_ index: Int) {
        
        let alert = UIAlertView.init(title:"Alert" , message: titleArray[index], delegate: self, cancelButtonTitle: "OK")
        alert.show()

    }

    func callBackFirstDisplayView(_ imageView: UIImageView, _ url: [String], _ index: Int) {
        
        imageView.kf.setImage(with: URL(string: url[index]), placeholder: UIImage.init(named: "defaultImage"), options: [.transition(.fade(1))])
        
    }
    
    func startAutoScroll() {
        carouselView.startScrollImageView()
        
    }
    
    func stopAutoScroll() {
        carouselView.stopScrollImageView()
    }
    
}

