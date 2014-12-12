pubApp.controller("EditController", ["$scope", "res", function($scope, res) {
	$scope.post = backend.post;
	$scope.cats = backend.cats;
	$scope.saving = false;

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