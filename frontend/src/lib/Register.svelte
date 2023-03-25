<script>
	import { useForm, Hint, HintGroup, validators, required, email } from "svelte-use-form";
	
	const form = useForm();
	
	const requiredMessage = "This field is required";
</script>

<main>
	<form use:form>
		<h1>
			Registration
		</h1>

        <label for="name">Name</label>
		<input type="text" name="name"  />

		<label for="email">Email</label>
		<input type="email" name="email" use:validators={[required, email]} />
		<HintGroup for="email">
			<Hint on="required">{requiredMessage}</Hint>
			<Hint on="email" hideWhenRequired>This must be a valid email</Hint>	
		</HintGroup>



		<button disabled={!$form.valid} on:click|preventDefault>
			Submit
		</button>
	</form>
</main>
	

<style>
	main {
		display: flex;
		justify-content: space-around;
        position: absolute;
		right: 2vw;
		top: 15vh;
		color: #b33636;
		background-color: white;
		border-color: #b33636;
		border-style: solid;
		border-width: 2px;
		border-radius: 20px;
		padding: 1vw;
		width: 450px;
	}

	form {
		display: flex;
		flex-direction: column;
	}

	@media (max-width: 768px) {
		main {
			width: 92vw;
		}
 	}
</style>