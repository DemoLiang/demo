//
//  ViewController.m
//  gpsmod
//
//  Created by nickdai on 2018/1/7.
//  Copyright © 2018年 nickdai. All rights reserved.
//

#import "ViewController.h"
#import "gpsmod-Swift.h"
#import "transform.h"
@interface ViewController ()

@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    double lat = 22.5432883286;
    double lon = 113.9589285851;
    
    double latChange;
    double lonChange;
    gcj2wgs(lat, lon, &latChange, &lonChange);
    
    
    NSLog(@"lat:%f lon:%f", latChange, lonChange);

}


- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}


@end
