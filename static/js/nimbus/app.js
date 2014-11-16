
(function(window, document, location, navigator, angular, undefined) {
    'use strict';

    angular.module( 'nimbus', [
        'controller.main',
        'ui.router',
        'controller.cai'
    ])

    .config(function($stateProvider, $urlRouterProvider){
        $urlRouterProvider.otherwise("/");
        $stateProvider
        .state('main',{
            url:"/",
            views:{
                'current@':{
                    templateUrl:"static/js/nimbus/controllers/main/main.html",
                    controller:"main"
                }
            }
        })
        $stateProvider
        .state('cai',{
            url:"/cai",
            views:{
                'current@':{
                    templateUrl:"static/js/nimbus/controllers/cai/cai.html",
                    controller:"cai"
                },
                'header@':{
                    templateUrl:"static/js/nimbus/controllers/main/header.html",
                }
            }
        });
    });


;}(window, document, location, navigator, angular, undefined));


