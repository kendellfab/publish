pubApp.controller("EditController", ["$scope", "res", function($scope, res) {
	$scope.post = backend.post;
	$scope.cats = backend.cats;

	$scope.savePost = function() {
		res.post.update($scope.post);
	};

	$scope.publishClick = function() {
		res.post.publish($scope.post);
	};
}]);