<template>
  <div class="box">
    <div class="block">
      <img src="@/assets/flaningo.jpeg" alt="Menuplanner Logo" width="32" class="mr-4"/>
      <span class="title has-text-grey">Menuplanner</span>
    </div>
    <b-field v-if="hasGoogleAuth">
      <AuthGoogle/>
    </b-field>
    <b-collapse :open.sync="showForm" aria-id="loginForm">
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
          <AuthEmailSignIn :init-email="email" :init-password="password"/>
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
  props: {
    email: {
      type: String,
      default: '',
    },
    password: {
      type: String,
      default: '',
    },
    hasGoogleAuth: {
      type: Boolean,
      default: false,
    },
    isExpanded: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    showForm: {
      get() {
        return !this.hasGoogleAuth || this.isExpanded;
      },
      set(newValue) {
        this.isExpanded = newValue;
      },
    },
  },
};
</script>

<style scoped>

</style>
