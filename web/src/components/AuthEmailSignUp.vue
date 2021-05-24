<template>
  <form @submit.prevent="signUp">
    <b-notification
      v-if="error"
      type="is-danger is-light"
      aria-close-label="Close error message"
      role="alert">
      {{ error }}
    </b-notification>
    <b-field label="Email">
      <b-input
        v-model="email"
        icon="envelope"
        type="email"
        placeholder="e.g. natalia@flaningo.ru"
        required
      />
    </b-field>
    <b-field label="Password" class="mb-4">
      <b-input
        v-model="password"
        type="password"
        icon="lock"
        placeholder="*******"
        required
        password-reveal
      />
    </b-field>
    <b-field label="Confirm password" class="mb-4">
      <b-input
        ref="passwordConfirm"
        v-model="passwordConfirm"
        type="password"
        icon="check"
        placeholder="*******"
        required
        password-reveal
      />
    </b-field>
    <b-button class="is-primary" label="Register" native-type="submit"/>
  </form>
</template>

<script>
export default {
  name: 'AuthEmailSignUp',
  data() {
    return {
      email: '',
      password: '',
      passwordConfirm: '',
      error: '',
    };
  },
  methods: {
    async signUp() {
      if (this.password !== this.passwordConfirm) {
        this.$refs.passwordConfirm.setValidity('is-danger', 'Passwords do not match');
        return;
      }
      const { email, password, passwordConfirm } = this;
      try {
        await this.$store.dispatch('signUp', { email, password, passwordConfirm });
        await this.$router.push('/');
      } catch (e) {
        this.error = e.response.data;
      }
    },
  },
};
</script>

<style scoped>

</style>
