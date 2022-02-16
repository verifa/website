<script lang="ts">
	import { seo } from '$lib/seo/store';

	seo.reset();
	$seo.title = 'Contact us';
</script>

<section>
	<h1>Get in touch.</h1>
</section>
<section class="max-w-4xl">
	<form
		id="contact-form"
		action="https://usebasin.com/f/ebc4ce09edb4"
		method="POST"
		class="flex flex-col gap-y-8"
	>
		<!-- Honeypot field to avoid spamming  -->
		<input type="hidden" name="_tunaboat" />
		<div>
			<label for="name" class="block">Name</label>
			<input type="text" name="name" id="name" autocomplete="name" required />
		</div>
		<div>
			<label for="email" class="block">Work Email</label>
			<input type="email" name="email" id="email" autocomplete="email" required />
		</div>
		<div>
			<label for="message" class="block">Message</label>
			<textarea id="message" name="message" rows="4" required />
		</div>
		<div class="flex gap-x-4 items-center">
			<input
				type="checkbox"
				id="newsletter"
				name="newsletter"
				class="border-[3px] border-v-black focus:ring-0 text-v-lilac"
			/>
			<label for="newsletter" class="block mb-0">Subscribe to our monthly newsletter</label>
		</div>
		<div>
			<button type="submit">Submit</button>
		</div>
		<p>
			By submitting this form you agree to our <a href="/privacy" class="underline"
				>Privacy Policy</a
			>
		</p>
	</form>
</section>
