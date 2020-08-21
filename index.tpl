<!DOCTYPE html>
<style>
	.regular {
		color: #00ffff;
	}
	.irregular {
		color: #b22222;
	}
</style>

<html>
<body>
	<form action="/" method="post">
    <p> Please Input Text <br>
        <textarea name="comment" cols="40" rows="10">{{.InputText}}</textarea>
    </p> 
	<div>
		<button> Colorize </button>
	</div>
	</form>
	<p> Colorized Text <br>
		{{.ColorizedText}}
	</p>
</body>
</html>
