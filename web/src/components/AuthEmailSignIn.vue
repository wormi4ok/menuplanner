<template>
  <form class="block" @submit.prevent="signIn">
    <b-field label="Email">
      <b-input
        v-model="email"
        icon="envelope"
        type="email"
        placeholder="e.g. natalia@flaningo.ru"
        required
      />
    </b-field>
    <b-field label="Password">
      <b-input
        v-model="password"
        type="password"
        icon="lock"
        placeholder="*******"
        required
        password-reveal
      />
    </b-field>
    <b-button class="is-primary" label="Continue" native-type="submit"/>
  </form>
</template>

<script>
export default {
  name: 'AuthEmailSignIn',
  data() {
    return {
      email: '',
      password: '',
    };
  },
  methods: {
    signIn() {
      // submit form here
      console.log(this.email);
    },
  },
};
</script>

<style scoped>

</style>
