<template>
  <div class="container">
    <UserDataForm
      target="http://yanap.docker.com:3000/login"
      :onResult="onSignIn"
      :onError="onSiginInFailed"
      buttonText="Sign In"
    />
  </div>
</template>

<script>
import UserDataForm from "~/components/UserDataForm.vue";
export default {
  components: {
    UserDataForm,
  },

  methods: {
    onSignIn(json) {
      var resultCode = json["result_code"];
      if (resultCode !== 0) {
        switch (resultCode) {
          case 1:
            alert("User is not register");
            break;
          case -1:
            alert("Fatal Error.");
            break;
        }
        return;
      }

      alert("Sign in Success!");
    },
    onSignInFailed(err) {
      alert("Sign in Failed. Error:" + err);
    },
  },
};
</script>

<style>
</style>
