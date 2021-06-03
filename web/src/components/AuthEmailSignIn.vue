<template>
  <form class="block" @submit.prevent="signIn">
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
    <b-field>
      <div class="level">
        <div class="level-left">
          <div class="level-item">
            <b-button class="is-primary" label="Continue" native-type="submit"/>
          </div>
        </div>
        <div class="level-right">
          <div class="level-item">
            <b-tooltip
              type="is-light" multilined position="is-right" :delay="300"
              label="Please, contact the administrator of this website to reset your password."
            >
              <b-button tag="a" label="Forgot password?" size="is-small" type="is-ghost"/>
            </b-tooltip>
          </div>
        </div>
      </div>
    </b-field>
  </form>
</template>

<script>
export default {
  name: 'AuthEmailSignIn',
  props: {
    initEmail: {
      type: String,
      default: '',
    },
    initPassword: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      userEmail: '',
      userPassword: '',
      error: '',
    };
  },
  computed: {
    email: {
      get() {
        return this.userEmail || this.initEmail;
      },
      set(newValue) {
        this.userEmail = newValue;
      },
    },
    password: {
      get() {
        return this.userPassword || this.initPassword;
      },
      set(newValue) {
        this.userPassword = newValue;
      },
    },
  },
  methods: {
    async signIn() {
      this.error = '';
      const { email, password } = this;
      try {
        await this.$store.dispatch('logIn', { email, password });
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
