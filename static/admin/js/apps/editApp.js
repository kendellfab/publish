pubApp.controller("EditController", ["$scope", "res", function($scope, res) {
	$scope.post = backend.post;

	$scope.savePost = function() {
		res.post.update($scope.post);
	};
}]);