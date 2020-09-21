<template>
  <div class="container">
    <h1>Sign In</h1>
    <UserDataForm
      target="http://yanap.docker.com:3000/login"
      :onResult="onSignIn"
      :onError="onSignInFailed"
      buttonText="Sign In"
    />
    <br />
    <button>
      <nuxt-link to="/signup" class="signup_button_link">Sign Up</nuxt-link>
    </button>
    <br />

    <nuxt-link to="/mypage">マイページへ（デバッグ用）</nuxt-link>
    <br />
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

      alert("Sign in Success! Token:" + json["token"]);
    },
    onSignInFailed(err) {
      alert("Sign in Failed. Error:" + err);
    },
  },
};
</script>

<style>
.signup_button_link {
  text-decoration: none;
  color: black;
}
</style>
