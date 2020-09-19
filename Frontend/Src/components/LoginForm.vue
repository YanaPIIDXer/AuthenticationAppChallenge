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
import axios from "axios";

export default {
  methods: {
    async login() {
      const res = await axios.get("http://yanap.docker.com");
      alert(res.data);
    },
  },
};
</script>

<style>
</style>
