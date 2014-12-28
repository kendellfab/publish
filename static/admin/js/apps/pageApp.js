pubApp.controller('PageController', ['$scope', 'res', function($scope, res){
	$scope.page = backend.page;
	$scope.saving = false;

	$scope.savePage = function() {
		$scope.saving = true;
		res.page.update($scope.page, success, error);
	}

	$scope.publishClick = function() {
		$scope.saving = true;
		res.page.update($scope.page, success, error);
	}

	function success(result) {
		$scope.saving = false;
	}

	function error(result) {
		$scope.saving = false;
	}
}]);