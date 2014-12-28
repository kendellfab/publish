pubApp.controller("SeriesController", ["$scope", "res", function($scope, res) {
	$scope.series = backend.series;
	$scope.saving = false;

	$scope.saveSeries = function() {
		$scope.saving = true;
		res.series.update($scope.series, success, error);
	}

	$scope.removePost = function(index) {
		alert(index);
	}

	function success(result) {
		$scope.saving = false;
	}

	function error(result) {
		$scope.saving = false;
	}
}]);