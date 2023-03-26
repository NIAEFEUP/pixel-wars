<script>
	import { useForm, Hint, HintGroup, validators, required, email } from "svelte-use-form";
	
	const form = useForm();
	
	const requiredMessage = "This field is required";
</script>

<main>
	<form use:form>

    <label for="name">Name</label>
		<input type="text" name="name"  use:validators={[required]} />
    <HintGroup for="name">
			<Hint on="required">{requiredMessage}</Hint>
		</HintGroup>

		<label for="email">Email</label>
		<input type="email" name="email" use:validators={[required, email]} />
		<HintGroup for="email">
			<Hint on="required">{requiredMessage}</Hint>
			<Hint on="email" hideWhenRequired>This must be a valid email</Hint>	
		</HintGroup>

    <label for="image">Upload a picture:</label>
    <input
      accept=".png, .jpeg, .jpg"
      type="file"
    />


		<button class="submit" disabled={!$form.valid} on:click|preventDefault>
			Submit
		</button>
	</form>
</main>
	

<style>
	main {
		display: flex;
		flex-direction: column;
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

	.submit {
		color: white;
		background-color: #b33636;
		cursor: pointer;
    border-radius: 10px;
		margin-top: 10px;
		outline: none;
		font-size: 1.5vh;
		font-weight: bolder;
    padding: 10px;
		position: relative;
    width: fit-content;
		text-align: center;
		text-decoration: none;
		vertical-align: baseline;
		touch-action: manipulation;
    align-self: center;
	}

  .submit:hover {
    background-color: #b33636ea;
    transform: scale(101%);
  }

  .submit:disabled {
    color: #ffffff96;
    background-color: #c4c4c496;
    cursor: unset;
  }

  label {
    margin-top: 15px;
  }

  input {
    border-width: 0px 0px 2px 0px;
    border-color: #b33636;
    margin-bottom: 10px;
  }

	@media (max-width: 768px) {
		main {
			width: 92vw;
		}
 	}
</style>