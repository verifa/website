<svelte:head>
	<!-- Add MailChimp integration for newsletter submission -->
	<script id="mcjs">
		!(function (c, h, i, m, p) {
			(m = c.createElement(h)),
				(p = c.getElementsByTagName(h)[0]),
				(m.async = 1),
				(m.src = i),
				p.parentNode.insertBefore(m, p);
		})(
			document,
			'script',
			'https://chimpstatic.com/mcjs-connected/js/users/8f84cc5c0851a90e124b335d3/b62311a48b6591e2e2820e0aa.js'
		);
	</script>
</svelte:head>

<div class="bg-v-black py-12">
	<div class="mx-auto max-w-7xl px-8 sm:px-16">
		<h2 class="text-v-white">Sign up for our monthly newsletter.</h2>

		<div id="mc_embed_signup" class="">
			<form
				action="https://verifa.us19.list-manage.com/subscribe/post?u=8f84cc5c0851a90e124b335d3&amp;id=a8220a54a9"
				method="post"
				id="mc-embedded-subscribe-form"
				name="mc-embedded-subscribe-form"
				class="validate"
				target="_blank"
				novalidate
			>
				<div
					id="mc_embed_signup_scroll"
					class="flex flex-col md:items-center gap-y-4 md:flex-row md:gap-x-8"
				>
					<div class="w-56 md:w-1/3">
						<label for="mce-EMAIL" class="sr-only">Email</label>
						<input
							type="email"
							name="EMAIL"
							class="email"
							id="mce-EMAIL"
							placeholder="you@example.com"
							required
						/>
					</div>
					<!-- real people should not fill this in and expect good things - do not remove this or risk form bot signups-->
					<div style="position: absolute; left: -5000px;" aria-hidden="true">
						<input
							type="text"
							name="b_8f84cc5c0851a90e124b335d3_a8220a54a9"
							tabindex="-1"
							value=""
						/>
					</div>
					<div>
						<button
							id="mc-embedded-subscribe"
							class="v-button bg-v-white text-v-black font-semibold"
							value="Subscribe">Submit</button
						>
					</div>
				</div>
			</form>
		</div>

		<span class="text-v-white mb-0 mt-4">
			By submitting this form you agree to our <a class="text-white" style="text-decoration: underline" href="/privacy"
				>Privacy Policy</a
			>
		</span>
	</div>
</div>
