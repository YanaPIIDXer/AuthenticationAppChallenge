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
      const res = await axios
        .get("http://yanap.docker.com:3000/login")
        .catch((err) => alert(err));
      if (!res || !res.data) {
        return;
      }
      alert("ResultCode:" + res.data["result_code"]);
    },
  },
};
</script>

<style>
</style>
