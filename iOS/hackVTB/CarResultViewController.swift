//
//  CarTableViewController.swift
//  hackVTB
//
//  Created by andarbek on 09.10.2020.
//

import UIKit
import Alamofire
import SwiftyJSON

class CarResultViewController: UIViewController, UIImagePickerControllerDelegate, UINavigationControllerDelegate, UITabBarControllerDelegate {
    
    var image : UIImageView?
    
    var car = Car()
    var jsonData : JSON?
    
    @IBOutlet weak var titleCar: UILabel?
    @IBOutlet weak var cardCar: UIView?
    @IBOutlet weak var jsonCar: UIImageView?
    @IBOutlet weak var price: UILabel?
    
    @IBOutlet weak var imageCar: UIImageView!
    @IBOutlet weak var activityView: UIActivityIndicatorView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        getPhoto()
        self.cardCar?.layer.cornerRadius = 8;
        self.cardCar?.layer.masksToBounds = true;
        self.imageCar.roundCornersForAspectFit(radius: 8.0)
        activityView.hidesWhenStopped = true
        activityView.color = #colorLiteral(red: 0, green: 0.1137531176, blue: 0.4273943305, alpha: 1)
        hiddenFunc(flag: true)
        let tap = UITapGestureRecognizer(target: self, action: #selector(self.handleTap(_:)))
        cardCar?.addGestureRecognizer(tap)
    }
    
    func hiddenFunc(flag: Bool){
        titleCar?.isHidden = flag
        cardCar?.isHidden = flag
        jsonCar?.isHidden = flag
        price?.isHidden = flag
    }
    
    @objc func handleTap(_ sender: UITapGestureRecognizer? = nil) {
        print("afghakjguhl iweuHF LieuwhfliwEHFLIuwehfliuwh")

        self.performSegue(withIdentifier: "showCarInfo", sender: self.cardCar)
    }
    
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?)
    {
        if segue.identifier == "showCarInfo" {
            if let nextVC = segue.destination as? CarInfoViewController {
                nextVC.jsonData = self.jsonData
                
            }
        }
    }
    
    
    func convertImageToBase64String (img: UIImage) -> String {
        return img.jpegData(compressionQuality: 1)?.base64EncodedString() ?? ""
    }
    
    func getPhoto(){
        self.imageCar.image = Photo.photo
        
        
        let parameters: [String: String] = [
            "content": self.convertImageToBase64String(img: (self.imageCar.image?.scaleImage(toSize: CGSize(width: 400, height: 400)))!)
        ]
        activityView?.startAnimating()
        AF.request("http://34.123.226.60:8081/recognition", method: .post, parameters: parameters as Parameters, encoding: JSONEncoding.default).responseJSON(completionHandler: { [self] response in
            if response.response?.statusCode == 404 {
                print("error 404")
            }
            else if response.response?.statusCode == 200 {
                do{
                    let json = try JSON(data: response.data!)
                    self.jsonData = json
                    let parameters: [String: Any] = [
                        "UserID": 1,
                        "JsonText" : jsonData?.description
                    ]
                    
                    let headers: HTTPHeaders = [
                        "x-access-token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIk5hbWUiOiIiLCJFbWFpbCI6InJvbWFhY3RhcG92QGdtYWlsLmNvbSIsImV4cCI6MTYwODM5Njg3Mn0.ZcQXzPSXv17hssstgw_oH-hsyzlidVx09kPQopig5BA"
                        ]
                    
                    AF.request("http://34.123.226.60:8080/auth/history", method: .post, parameters: parameters as Parameters, encoding: JSONEncoding.default, headers: headers).responseJSON(completionHandler: { [weak self] response in
                        print(response)
                    }).resume()
                    car.title = json["Cars"][0]["Title"].string
                    car.info = String(decoding: response.data!, as: UTF8.self)
                    car.price = json["Cars"][0]["prettyPrice"].string
                    car.image = json["Cars"][0]["photo"].string
                    print(car)
                    downloadImage(from: (car.image)!)
                    DispatchQueue.main.async {
                        self.titleCar?.text = car.title
                        self.price?.text = car.price
                        self.activityView.stopAnimating()
                        self.hiddenFunc(flag: false)
                    }
                } catch{
                    
                }
            }
            else{
                print("error")
            }
        }).resume()
    }
    
    
    
    
    func downloadImage(from url: String) {
        
        guard let url = URL(string: url) else {
            return
        }
        getData(from: url) { data, response, error in
            guard let data = data, error == nil else { return }
            DispatchQueue.main.async { [weak self] in
                self?.jsonCar?.image = UIImage(data: data)
            }
        }
    }
    func getData(from url: URL, completion: @escaping (Data?, URLResponse?, Error?) -> ()) {
        URLSession.shared.dataTask(with: url, completionHandler: completion).resume()
    }
    
}


