<template>
  <b-navbar-dropdown>
    <template #label>
      <strong>{{ displayName }}</strong>
    </template>
    <b-navbar-item @click="onLogout">
      Logout
    </b-navbar-item>
  </b-navbar-dropdown>
</template>

<script>
import { mapActions, mapState } from 'pinia';
import { useUserStore } from '@/stores/user';

export default {
  name: 'NavbarUserMenu',
  computed: {
    ...mapState(useUserStore, ['currentUser']),
    displayName() {
      return this.currentUser.name || this.currentUser.email;
    },
  },
  methods: {
    ...mapActions(useUserStore, ['logOut']),
    onLogout() {
      this.logOut();
      this.$router.push({ name: 'Login' });
    },
  },
};
</script>

<style scoped>

</style>
