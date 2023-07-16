<script>
	import { fade } from 'svelte/transition';
</script>

<div
	class="hero min-h-screen"
	style="background-image: url(https://boluswizard.io/assets/loginpage/bg.gif);"
>
	<div class="hero-overlay" />
	<div
		class="hero-content text-center text-neutral-content"
	>
		<div class="max-w-md text-white">
			<h1 class="mb-5 text-5xl font-bold">Welcome to BolusWizard.io!</h1>
			<p class="mb-5">
				Don't have an expensive insulin pump? No problem! BolusWizard calculates the amount of
				insulin you need for meals and corrections! Just enter your Insulin sensitivity factor,
				Insulin duration and BG targets and you're ready to go!
			</p>
			<a class="btn btn-primary" href="/signin">Sign In</a>
			<a class="btn btn-secondary ml-4" href="/signup">Sign Up</a>
		</div>
	</div>
</div>
