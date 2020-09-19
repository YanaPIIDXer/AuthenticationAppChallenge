<template>
  <div class="container">
    <h1>Sign Up</h1>
    <UserDataForm
      target="http://yanap.docker.com:3000/signup"
      :onResult="onSignUp"
      :onError="onSignUpFailed"
      buttonText="Sign Up"
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
    onSignUp(json) {},
    onSignUpFailed(err) {
      alert("Sign up Failed. Error:" + err);
    },
  },
};
</script>

<style>
</style>
