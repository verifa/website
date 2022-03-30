<div class="max-w-4xl">
	<form
		action="https://usebasin.com/f/9f1e7bc2e780"
		method="POST"
		enctype="multipart/form-data"
		id="contact-form"
		class="flex flex-col gap-y-8"
	>
		<!-- Honeypot field to avoid spamming  -->
		<input type="hidden" name="_tunaboat" />
		<div>
			<label for="name" class="block">Name*</label>
			<input type="text" name="name" id="name" autocomplete="name" required />
		</div>
		<div>
			<label for="email" class="block">Email*</label>
			<input type="email" name="email" id="email" required />
		</div>
		<div>
			<label for="message" class="block">Tell us about yourself*</label>
			<textarea id="message" name="message" rows="4" required />
		</div>
		<div class="flex flex-col gap-y-8 md:flex-row md:gap-x-4">
			<div class="flex-1">
				<label for="linkedin" class="block">LinkedIn</label>
				<input
					type="url"
					name="linkedin"
					id="linkedin"
					placeholder="https://www.linkedin.com/in/<profile>"
					pattern="https://(www.)?linkedin.com/.*"
				/>
			</div>
			<div class="flex-1">
				<label for="github" class="block">GitHub</label>
				<input
					type="url"
					name="github"
					id="github"
					placeholder="https://github.com/<profile>"
					pattern="https://(www.)?github.com/.*"
				/>
			</div>
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
			<button class="bg-v-black" type="submit">Submit</button>
		</div>
		<p>
			By submitting this form you agree to our <a href="/privacy" class="underline"
				>Privacy Policy</a
			>
		</p>
	</form>
</div>
