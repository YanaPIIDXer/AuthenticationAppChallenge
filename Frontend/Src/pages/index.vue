<template>
  <div class="container">
    <UserDataForm
      target="http://yanap.docker.com:3000/login"
      :onResult="onSignIn"
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
      alert("ResultCode:" + json["result_code"]);
    },
  },
};
</script>

<style>
</style>
