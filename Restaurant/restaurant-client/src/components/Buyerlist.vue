<template>
  <v-container>
    <h1>Buyer List</h1>
    <div v-if="errored">An error occurred while loading this page.</div>
    <div v-else-if="!loading">
      <v-list>
        <template v-for="(buyer, index) in buyers.slice((this.page - 1) * rowsPerPage, page * rowsPerPage)">
          <v-list-item :key="index" class="ma-0 pa-0">
            <v-list-item-content class="ma-0 pa-0">
              <v-container class="ma-0 pa-0">
                <v-row
                  align-content="start"
                  no-gutters
                  class="ma-0 pa-0"
                  align="center"
                  justify="center"
                >
                  <v-col>
                    <v-list-item-action>
                      <v-btn depressed color="primary" min-width="150px" v-on:click="goToHistory(buyer.id)">
                        ID: {{ buyer.id }}
                      </v-btn>
                    </v-list-item-action>
                  </v-col>
                  <v-col><strong>Name: </strong>{{ buyer.name }}</v-col>
                  <v-col><strong>Age: </strong>{{ buyer.age }}</v-col>
                </v-row>
              </v-container>

              <v-divider></v-divider>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-list>

      <div class="text-center">
        <v-pagination
          v-model="page"
          :length="Math.ceil(buyers.length / rowsPerPage)"
          :total-visible="7"
        ></v-pagination>
      </div>
    </div>
    <div v-else><h2>Loading...</h2></div>
  </v-container>
</template>

<script>

import axios from 'axios'

export default {
  name: 'Buyerlist',
  data () {
    return {
      rowsPerPage: 30,
      page: 1,
      buyers: [],
      loading: true,
      errored: false
    }
  },

  mounted () {
    axios
      .get(' http://localhost:5000/api/buyers?first=2147483647&offset=0')
      .then(response => {
        this.buyers = response.data
      })
      .catch(error => {
        console.log(error)
        this.buyers = []
        this.errored = true
      })
      .finally(() => {
        this.loading = false
      })
  },
  methods: {
    goToHistory (id) {
      this.$router.push('/buyer/' + id)
    }
  }
}
</script>
