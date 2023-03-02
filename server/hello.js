angular.module('zombeez', [])
    .controller('Today', function ($scope, dataService) {
        dataService.getData('today').then(function (data) {
            $scope.today = data;
        });

        var setTimeOut = setInterval(function () {
            dataService.getData('today').then(function (data) {
                $scope.today = data;
            });
        }, 300000);
    }).controller('Yesterday', function ($scope, dataService) {
        dataService.getData('yesterday').then(function (data) {
            $scope.yesterday = data;
        });
    }).controller('Weekly', function ($scope, dataService) {
        dataService.getData('weekly').then(function (data) {
            $scope.weekly = data;
        });

        var setTimeOut = setInterval(function () {
            dataService.getData('weekly').then(function (data) {
                $scope.weekly = data;
            });
        }, 300000);
    }).controller('Tally', function ($scope, dataService) {
        dataService.getData('tally').then(function (data) {
            $scope.tally = data;
        });
    }).controller('Coolpoints', function ($scope, dataService) {
        dataService.getData('coolpoints').then(function (data) {
            $scope.coolpoints = data;
        });
    }).controller('Custom420', function ($scope, dataService) {
        dataService.getData('custom420').then(function (data) {
            $scope.custom420 = data;
        });
    }).controller('Custom69', function ($scope, dataService) {
        dataService.getData('custom69').then(function (data) {
            $scope.custom69 = data;
        });
    }).controller('Custom007', function ($scope, dataService) {
        dataService.getData('custom007').then(function (data) {
            $scope.custom007 = data;
        });
    }).controller('High', function ($scope, dataService) {
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
