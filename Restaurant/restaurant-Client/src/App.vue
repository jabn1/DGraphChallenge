<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
    <router-link to="/home">
    <h1 style="color: white;" >Restaurant</h1>
    </router-link>

    <v-btn to="/buyers" color="secondary" text class="pl-10">
      Buyer List
    </v-btn>
    <v-btn to="/buyer" color="secondary" text>
      Buyer History
    </v-btn>
    <v-btn to="/load" color="secondary" text>
      Date Loader
    </v-btn>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App'
}
</script>
