//
//  CustomTabBarController.swift
//  hackVTB
//
//  Created by andarbek on 10.10.2020.
//

import UIKit

class CustomTabBarController: UITabBarController, UIImagePickerControllerDelegate, UINavigationControllerDelegate, UITabBarControllerDelegate{
    
    
    
    var image: UIImage?
    let menuButton = UIButton(frame: CGRect.zero)
    
    override func viewDidLoad() {
        super.viewDidLoad()
        self.delegate = self
        self.view.bringSubviewToFront(self.tabBar)
        self.setupMiddleButton()
    }
    
    func setupMiddleButton() {
        let image = UIImage(named: "camera0")?.scaleImage(toSize: CGSize(width: 50, height: 50))
        let numberOfItems = CGFloat(tabBar.items!.count)
        let tabBarItemSize = CGSize(width: tabBar.frame.width / numberOfItems, height: tabBar.frame.height)
        menuButton.frame = CGRect(x: 0, y: 0, width: 80, height: 80)
        var menuButtonFrame = menuButton.frame
        menuButtonFrame.origin.y = self.view.bounds.height - menuButtonFrame.height - self.view.safeAreaInsets.bottom
        menuButtonFrame.origin.x = self.view.bounds.width/2 - menuButtonFrame.size.width/2
        menuButton.setImage(image, for: .normal)
        menuButton.frame = menuButtonFrame
        menuButton.layer.cornerRadius = 0.5 * menuButton.bounds.size.width
        menuButton.clipsToBounds = true
        menuButton.backgroundColor = #colorLiteral(red: 0.08361851424, green: 0.2466595173, blue: 0.4754746556, alpha: 1)
        
        self.view.addSubview(menuButton)
        self.view.layoutIfNeeded()
        
        menuButton.addTarget(self, action: #selector(didTouchCenterButton(_:)), for: .touchUpInside)
    }
    
    override func viewDidLayoutSubviews() {
        super.viewDidLayoutSubviews()
        menuButton.frame.origin.y = self.view.bounds.height - menuButton.frame.height - self.view.safeAreaInsets.bottom
    }
    
    @objc private func didTouchCenterButton(_ sender: UIButton) {
        print("hi")
        cameraWork()
    }
    
 
    
    func cameraWork() {
        let imagePicker = UIImagePickerController()
        imagePicker.delegate = self
        let alert = UIAlertController(title: "Image Selection", message: "from source", preferredStyle: .actionSheet)
        
        alert.addAction(UIAlertAction(title: "Camera", style: .default, handler: { (action: UIAlertAction) in
            imagePicker.sourceType = .camera
            imagePicker.isEditing = false
            self.present(imagePicker, animated: true, completion: nil)
            
        }))
        alert.addAction(UIAlertAction(title: "Photo library", style: .default, handler: {(action: UIAlertAction) in
            imagePicker.sourceType = .photoLibrary
            imagePicker.isEditing = false
            imagePicker.allowsEditing = true
            self.present(imagePicker, animated: true, completion: nil)
        }))
        alert.addAction(UIAlertAction(title: "Cancel", style: .cancel, handler: nil))
        self.present(alert, animated: true, completion: nil)
    }
    
    
    func imagePickerController(_ picker: UIImagePickerController, didFinishPickingMediaWithInfo info: [UIImagePickerController.InfoKey : Any]) {
        guard let image = info[UIImagePickerController.InfoKey.originalImage] as? UIImage else {
            print("No image found")
            return
        }
        picker.dismiss(animated: true, completion: nil)
        Photo.photo = image
        print(Photo.photo)
        segue()
    }
    
    func segue(){
        if let tabViewController = storyboard?.instantiateViewController(withIdentifier: "carList") as? CarResultViewController {
            present(tabViewController, animated: true, completion: nil)
        }
       // self.performSegue(withIdentifier: "showCars", sender:self)
    }
    
//        override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
//            if segue.identifier == "showCars"{
//                if let nextView = segue.destination as? CarListViewController{
//                    if let image = self.image{
//                        nextView.imageCar = UIImageView(image: image)
//                    }
//                    print("hello")
//                }
//            }
//        }
    
    
    
    func imagePickerControllerDidCancel(_ picker: UIImagePickerController) {
        picker.dismiss(animated: true, completion: nil)
    }
    
    
}


