//
//  HistoryTableViewCell.swift
//  hackVTB
//
//  Created by andarbek on 11.10.2020.
//

import UIKit

class HistoryTableViewCell: UITableViewCell {

    @IBOutlet weak var carImageView: UIImageView!
    @IBOutlet weak var carName: UILabel!
    @IBOutlet weak var carPrice: UILabel!
    
    override func awakeFromNib() {
        super.awakeFromNib()
        self.preservesSuperviewLayoutMargins = false
        self.separatorInset = UIEdgeInsets.zero
        self.layoutMargins = UIEdgeInsets.zero
    }
    

    override func setSelected(_ selected: Bool, animated: Bool) {
        super.setSelected(selected, animated: animated)

    }
}
