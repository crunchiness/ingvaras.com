(function(angular) {
    'use strict';
    var myApp = angular.module('apiApp', ['ngAudio']);

    myApp.controller('apiCtrl', ['$scope', 'ngAudio', function($scope, ngAudio) {
        $scope.text = 'programmatically';
        $scope.say = function () {
          var query = 'http://www.ingvaras.com/tts/en/' + $scope.text;
          var sound = ngAudio.load(query);
          sound.play();
        };
    }]);
})(window.angular);
