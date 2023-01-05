angular.module('zombeez', [])
    .controller("main", function($scope, dataService) {
        dataService.getData('today').then(function (data) {
            $scope.today = data;
        });
        dataService.getData('yesterday').then(function (data) {
            $scope.yesterday = data;
        });
        dataService.getData('weekly').then(function (data) {
            $scope.weekly = data;
        });
        dataService.getData('tally').then(function (data) {
            $scope.tally = data;
        });
        dataService.getData('yesterdcoolpointsay').then(function (data) {
            $scope.coolpoints = data;
        });
        dataService.getData('custom420').then(function (data) {
            $scope.custom420 = data;
        });
        dataService.getData('custom69').then(function (data) {
            $scope.custom69 = data;
        });
        dataService.getData('high').then(function (data) {
            $scope.high = data;
        });
    }).factory('dataService', function ($http) {
        return {
            getData: function (params) {
                return $http.get('https://butts69420.com/' + params).then(function (response) {
                    return response.data;
                })
            }
        };
    });