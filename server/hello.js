angular.module('zombeez', [])
    .controller('Today', function ($scope, dataService) {
        $scope.today = dataService.getData('today');
    }).controller('Yesterday', function ($scope, dataService) {
        $scope.yesterday = dataService.getData('yesterday');
    }).controller('Weekly', function ($scope, dataService) {
        $scope.weekly = dataService.getData('weekly');
    }).controller('Tally', function ($scope, dataService) {
        $scope.tally = dataService.getData('tally');
    }).controller('Coolpoints', function ($scope, dataService) {
        $scope.coolpoints = dataService.getData('coolpoints');
    }).controller('Custom420', function ($scope, dataService) {
        $scope.custom420 = dataService.getData('custom420');
    }).controller('Custom69', function ($scope, dataService) {
        $scope.custom69 = dataService.getData('custom69');
    }).controller('High', function ($scope, dataService) {
        $scope.high = dataService.getData('high');
    }).factory('dataService', function ($http) {
        return {
            getData: function (params) {
                return $http.get('https://butts69420.com/' + params);
            }
        };
    });