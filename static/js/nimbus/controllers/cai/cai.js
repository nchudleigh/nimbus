(function(window, document, location, navigator, angular, undefined) {
    'use strict';

    angular.module('controller.cai',[])

    .controller('cai',['$scope', function($scope) {
            $scope.start=function(){
                console.log('starting..');
            };
    }])

;}(window, document, location, navigator, angular, undefined));