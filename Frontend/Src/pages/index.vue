<template>
  <div class="container">
    <UserDataForm target="http://yanap.docker.com:3000/login" buttonText="Sign In" />
  </div>
</template>

<script>
import UserDataForm from "~/components/UserDataForm.vue";
export default {
  components: {
    UserDataForm,
  },
};
</script>

<style>
</style>
