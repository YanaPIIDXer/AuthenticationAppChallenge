<template>
  <div>
    <form method="POST">
      <input type="text" v-model="email" placeholder="Input your email address." />
      <br />
      <input type="password" v-model="password" placeholder="Input your password." />
      <br />
      <input type="button" v-on:click="login()" :value="buttonText" />
      <br />
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      email: "",
      password: "",
    };
  },
  props: {
    target: String,
    onResult: Function,
    onError: {
      type: Function,
      default: function (err) {},
    },
    buttonText: {
      type: String,
      default: "Submit",
    },
  },

  methods: {
    async login() {
      const res = await axios
        .post(this.target, {
          email: this.email,
          password: this.password,
        })
        .catch((err) => this.onError(err));
      if (!res || !res.data) {
        return;
      }
      this.onResult(res.data);
    },
  },
};
</script>

<style>
</style>
