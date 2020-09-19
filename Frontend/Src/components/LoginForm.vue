<template>
  <div>
    <form method="POST">
      <input type="text" name="email" placeholder="Input your email address." />
      <br />
      <input type="password" name="password" placeholder="Input your password." />
      <br />
      <input type="button" v-on:click="login()" value="Login" />
      <br />
    </form>
  </div>
</template>

<script>
export default {
  methods: {
    login: function (event) {
      alert("Login");
    },
  },
};
</script>

<style>
</style>
