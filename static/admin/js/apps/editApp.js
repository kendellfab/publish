pubApp.controller("EditController", ["$scope", "res", function($scope, res) {
	$scope.post = backend.post;
	$scope.cats = backend.cats;
	$scope.series = backend.series;
	$scope.saving = false;

	$scope.uploads = res.uploads.list();

	$scope.savePost = function() {
		$scope.saving = true;
		res.post.update($scope.post, success, error);
	};

	$scope.publishClick = function() {
		$scope.saving = true;
		res.post.publish($scope.post, success, error);
	};

	function success(result) {
		$scope.saving = false;
	}

	function error(result) {
		$scope.saving = false;
	}
}]);