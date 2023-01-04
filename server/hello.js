angular.module('zombeez', [])
.controller('Today', function($scope, $http) {
    $http.get('https://butts69420.com/today').
        then(function(response) {
            $scope.today = response.data;
        });
}).controller('Yesterday', function($scope, $http) {
    $http.get('https://butts69420.com/yesterday').
        then(function(response) {
            $scope.yesterday = response.data;
        });
}).controller('Weekly', function($scope, $http) {
    $http.get('https://butts69420.com/weekly').
        then(function(response) {
            $scope.weekly = response.data;
        });
}).controller('Tally', function($scope, $http) {
    $http.get('https://butts69420.com/tally').
        then(function(response) {
            $scope.tally = response.data;
        });
}).controller('Coolpoints', function($scope, $http) {
    $http.get('https://butts69420.com/coolpoints').
        then(function(response) {
            $scope.coolpoints = response.data;
        });    
}).controller('Custom420', function($scope, $http) {
    $http.get('https://butts69420.com/custom420420').
        then(function(response) {
            $scope.custom420 = response.data;
        });    
}).controller('Custom69', function($scope, $http) {
    $http.get('https://butts69420.com/custom696969').
        then(function(response) {
            $scope.custom69 = response.data;
        });    
}).controller('High', function($scope, $http) {
    $http.get('https://butts69420.com/high').
        then(function(response) {
            $scope.high = response.data;
        });    
        
    
});


