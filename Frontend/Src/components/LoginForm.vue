<template>
  <div>
    <form method="POST">
      <input type="text" v-model="email" placeholder="Input your email address." />
      <br />
      <input type="password" v-model="password" placeholder="Input your password." />
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
        .post("http://yanap.docker.com:3000/login", {
          email: this.email,
          password: this.password,
        })
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
