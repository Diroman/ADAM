//
//  SearchViewController.swift
//  hackVTB
//
//  Created by andarbek on 10.10.2020.
//

import UIKit
import Alamofire
import SwiftyJSON

class HistoryViewController: UIViewController {

    struct Post {
        var image: String
        var imageView: UIImage?
        var name: String
        var price: String
    }
    
    var postList = [Post]()
    
    @IBOutlet weak var tableView: UITableView!
    @IBOutlet weak var activityView: UIActivityIndicatorView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        self.activityView.hidesWhenStopped = true
        
        //        tableView.delegate = self
        //        tableView.dataSource = self
        //self.tableView.register(CustomTableViewCell.self, forCellReuseIdentifier: "customCell")
        
        getData()
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        activityView.color = #colorLiteral(red: 0, green: 0.1137531176, blue: 0.4273943305, alpha: 1)
    }
    
    func getData(){
        self.activityView.startAnimating()
        let headers: HTTPHeaders = [
            "x-access-token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIk5hbWUiOiIiLCJFbWFpbCI6InJvbWFhY3RhcG92QGdtYWlsLmNvbSIsImV4cCI6MTYwODM5Njg3Mn0.ZcQXzPSXv17hssstgw_oH-hsyzlidVx09kPQopig5BA"
        ]
        
        AF.request("http://34.123.226.60:8080/auth/history/user/1", method: .get, headers: headers).responseJSON(completionHandler: { [weak self] response in
            do{
                let json = try JSON(data: response.data!)
                
                var a = Set<String>()
                for i in json{
                    let str = i.1["JsonText"].string
                    let data = str!.data(using: .utf8)
                    let jsondata = try? JSON(data: data!)
                    print(jsondata)
                    if !a.contains(jsondata!["Cars"][0]["Title"].string!) {
                        self?.postList.append(Post(image: jsondata!["Cars"][0]["photo"].string!, name: jsondata!["Cars"][0]["Title"].string!, price: jsondata!["Cars"][0]["prettyPrice"].string!))
                        a.insert(jsondata!["Cars"][0]["Title"].string!)
                    }
                    DispatchQueue.global().sync {
                        for i in 0..<self!.postList.count{
                            self!.downloadImage(from: self!.postList[i].image, at: i)
                        }
                    }
                }
                
                DispatchQueue.main.async { [weak self] in
                    self!.tableView.delegate = self as! UITableViewDelegate
                    self!.tableView.dataSource = self as! UITableViewDataSource
                    self!.tableView.reloadData()
                    self!.activityView.stopAnimating()
                }
                
                
            } catch{
                print("df")
            }
            
        }).resume()
    }
    
    
    func downloadImage(from url: String, at index: Int) {
        
        guard let url = URL(string: url) else {
            return
        }
        getData(from: url) { [weak self] data, response, error in
            guard let data = data, error == nil else { return }
            self!.postList[index].imageView = UIImage(data: data)
        }
    }
    func getData(from url: URL, completion: @escaping (Data?, URLResponse?, Error?) -> ()) {
        URLSession.shared.dataTask(with: url, completionHandler: completion).resume()
    }
    
}

extension HistoryViewController:  UITableViewDelegate, UITableViewDataSource {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return postList.count
    }


    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = self.tableView.dequeueReusableCell(withIdentifier: "historyCell") as! HistoryTableViewCell
        cell.carImageView.image = postList[indexPath.row].imageView
        cell.carName.text = postList[indexPath.row].name
        cell.carPrice.text = postList[indexPath.row].price
        return cell
    }

        func tableView(_ tableView: UITableView, commit editingStyle: UITableViewCell.EditingStyle, forRowAt indexPath: IndexPath) {
            if editingStyle == .delete {
                self.postList.remove(at: indexPath.row)
                self.tableView.deleteRows(at: [indexPath], with: .fade)
             }
        }

}
