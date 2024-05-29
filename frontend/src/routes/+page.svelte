<script lang="ts">
  import img from "$lib/assets/images/afr.png";
  import logo from "$lib/assets/images/logo.png";
  import FormError from "$lib/components/badges/FormError.svelte";
  import Button from "$lib/components/ui/button/Button.svelte";
  import TextField from "$lib/components/ui/input/TextField.svelte";

  export let form;
</script>

<div
  class="flex flex-col space-y-2 bg-navColor h-screen p-8 0 rounded items-center justify-center"
>
  <div
    class="w-full md:w-[500px] space-y-3 bg-white p-6 border-gray-300 rounded shadow-sm"
  >
    <div class="flex flex-row w-full items-center justify-center">
      <img src={logo} class="h-28 w-28" alt="" />

      <div class="flex flex-col w-full items-start justify-start">
        <p class="font-bold text-xl">Fusion Market</p>
        <p class="text-start text-[13px] font-medium text-gray-600">
          Join <span class="text-primaryColor">Fusion Market</span>
          today and create your own store and deploy your website and start selling
          immediately
        </p>
      </div>
    </div>

    <!-- {#if data.err !== undefined}
        <FormError title={"Unable to unauthorize"} message={data.err} err={true}/>
    {/if} -->

    <form method="POST" action="?/login" class="w-full flex flex-col space-y-3">
      <TextField label="Email" name="email" placeholder="enter your email" />
      <TextField
        label="Password"
        name="password"
        placeholder="Password"
        type="password"
      />

      <Button label="Sign In With Email" icon="line-md:login" />
    </form>
    <div class="flex items-center justify-center space-x-2">
      <div class="h-px bg-gray-300 w-16" />
      <!-- Left horizontal line -->
      <p
        class="text-[12px] font-semibold text-gray-500 text-center m-2 px-4 rounded-md"
      >
        Or
      </p>
      <div class="h-px bg-gray-300 w-16" />
      <!-- Right horizontal line -->
    </div>
    <!-- <div class="flex flex-col w-full space-y-3">
      <form action="?/googleLogin" method="post" class="flex flex-col w-full">
        <Button
          label="Continue With Google"
          variant="outlined"
          icon="devicon:google"
        />
      </form>

      <form action="?/githubLogin" method="post" class="flex flex-col w-full">
        <Button
          label="Continue With Github"
          variant="outlined"
          icon="line-md:github-loop"
        />
      </form>
    </div> -->

    <p class="text-center text-[14px]">
      <span class="text-primaryColor">Forgot Password?</span> | Not a Member?
      <a href="/register">
        <span class="text-primaryColor">Register</span>
      </a>
    </p>

    {#if form?.error}
      <div class="h-1" />
      <FormError
        err={true}
        title="Authorization Failed"
        message={form?.error}
      />
    {/if}
  </div>
</div>
