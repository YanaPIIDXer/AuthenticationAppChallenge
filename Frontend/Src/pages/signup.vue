<template>
  <div class="container">
    <h1>Sign Up</h1>
    <UserDataForm
      target="http://yanap.docker.com:3000/register"
      :onResult="onSignUp"
      :onError="onSignUpFailed"
      buttonText="Sign Up"
      enableConfirmPassword="Yes"
    />
    <br />
    <nuxt-link to="/">Back</nuxt-link>
  </div>
</template>

<script>
import UserDataForm from "~/components/UserDataForm.vue";
export default {
  components: {
    UserDataForm,
  },

  methods: {
    onSignUp(json) {
      var resultCode = json["result_code"];
      if (resultCode != 0) {
        alert("Sign Up Failed. Error:" + json["error_message"]);
        return;
      }
      alert("Sign Up");
    },
    onSignUpFailed(err) {
      alert("Sign up Failed. Error:" + err);
    },
  },
};
</script>

<style>
</style>
