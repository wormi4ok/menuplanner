<template>
  <div class="box">
    <div class="block">
      <img src="@/assets/flaningo.jpeg" alt="Menuplanner Logo" width="32" class="mr-4"/>
      <span class="title has-text-grey">Menuplanner</span>
    </div>
    <b-field v-if="googleAuthEnabled">
      <AuthGoogle/>
    </b-field>
    <b-collapse :open.sync="registrationForm" aria-id="loginForm">
      <template #trigger="props">
        <div class="has-text-centered">
          <a v-if="!props.open" aria-controls="loginForm" class="icon-text">
            <b-icon icon="chevron-down"/>
            <span>Use Email</span>
          </a>
        </div>
      </template>
      <b-tabs position="is-centered" size="is-large">
        <b-tab-item label="Sign In">
          <AuthEmailSignIn/>
        </b-tab-item>
        <b-tab-item label="Sign Up">
          <AuthEmailSignUp/>
        </b-tab-item>
      </b-tabs>
    </b-collapse>
  </div>
</template>

<script>
import AuthGoogle from '@/components/AuthGoogle.vue';
import AuthEmailSignIn from '@/components/AuthEmailSignIn.vue';
import AuthEmailSignUp from '@/components/AuthEmailSignUp.vue';

export default {
  name: 'AuthForm',
  components: {
    AuthGoogle,
    AuthEmailSignIn,
    AuthEmailSignUp,
  },
  data() {
    return {
      showForm: false,
    };
  },
  computed: {
    googleAuthEnabled() {
      return !!window.config.MP_CLIENT_ID;
    },
    registrationForm() {
      return this.showForm || !this.googleAuthEnabled;
    },
  },
};
</script>

<style scoped>

</style>
