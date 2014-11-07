(function(window, document, location, navigator, angular, undefined) {
    'use strict';

    angular.module('directive.main',[])

    .directive('main',[function() {
        return {
            restrict : 'EA',
            templateUrl : '/static/js/nimbus/directives/main/main.html',
        };
    }])

;}(window, document, location, navigator, angular, undefined));