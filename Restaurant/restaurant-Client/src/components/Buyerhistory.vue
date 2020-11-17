<template>
  <v-container>
    <!-- Go to other buyer -->
    <form class="pa-0 ma-0">
      <v-row class="pa-0 ma-0">
        <v-col cols="3" class=" pb-0 pl-0" >
          <v-text-field
            label="Buyer ID"
            solo
            dense
            v-model="input"
          ></v-text-field>
        </v-col >
        <v-col class=" pb-0">
          <v-btn v-on:click="goToBuyer">
            Go
          </v-btn>
        </v-col>
      </v-row>

    </form>
    <!-- Buyer history -->
    <h1>Buyer History</h1>
    <div v-if="id === undefined">Please enter a buyer's ID</div>
    <div v-else-if="errored">
      There was an error loading this page.
    </div>
    <div v-else>
      <div v-if="loading">
        Loading...
      </div>
      <div v-else>
        <div v-if="buyerHistory.buyer === undefined">A buyer with ID: <strong>{{id}}</strong> is not in the records.</div>
        <div v-else >
<!-- Start of main content -->
    <!-- Buyer -->
          <h3>Buyer</h3>
          <v-container class=" lighten-5">
            <v-row no-gutters>
              <v-col cols="12" sm="4" >
                <v-card class="pa-2" outlined tile >
                  ID: {{id}}
                </v-card>
              </v-col>
              <v-col cols="12" sm="4" >
                <v-card class="pa-2" outlined tile >
                  Name: {{buyerHistory.buyer.name}}
                </v-card>
              </v-col>
              <v-col cols="12" sm="4" >
                <v-card class="pa-2" outlined tile >
                  Age: {{buyerHistory.buyer.age}}
                </v-card>
              </v-col>
            </v-row>
          </v-container>
    <!-- Purchase History -->
          <h3>Purchase History</h3>
            <v-row justify="center">
              <v-expansion-panels accordion>
                <v-expansion-panel
                  v-for="(item,i) in buyerHistory.transactions.slice((this.page - 1) * perPage, page * perPage)"
                  :key="i"
                >
                  <v-expansion-panel-header>
                    <!-- Transaction info -->
                    <h5>Transaction:</h5>
                      <v-container class=" lighten-5">
                        <v-row no-gutters>
                          <v-col cols="auto"  >
                            <v-card class="pa-2" outlined tile >
                              <strong>ID: </strong>{{item.id}}
                            </v-card>
                          </v-col>
                          <v-col cols="auto"  >
                            <v-card class="pa-2" outlined tile >
                              <strong>IP: </strong>{{item.ip}}
                            </v-card>
                          </v-col>
                          <v-col cols="auto"  >
                            <v-card class="pa-2" outlined tile >
                              <strong>Device: </strong>{{item.device}}
                            </v-card>
                          </v-col>
                          <v-col cols="auto"  >
                            <v-card class="pa-2" outlined tile >
                              <strong>Date: </strong>{{item.date}}
                            </v-card>
                          </v-col>
                        </v-row>
                      </v-container>
                  </v-expansion-panel-header>
                  <v-expansion-panel-content>
                    <!-- Product list -->
                    <h5>Products:</h5>
                    <v-simple-table>
                      <template v-slot:default>
                        <thead>
                          <tr>
                            <th class="text-left">
                              ID
                            </th>
                            <th class="text-left">
                              Name
                            </th>
                            <th class="text-left">
                              Price
                            </th>
                          </tr>
                        </thead>
                        <tbody>
                          <tr
                            v-for="(product,i) in item.products"
                            :key="i"
                          >
                            <td>{{ product.id }}</td>
                            <td>{{ product.name }}</td>
                            <td>${{ product.price/100 }}</td>
                          </tr>
                        </tbody>
                      </template>
                    </v-simple-table>

                  </v-expansion-panel-content>
                </v-expansion-panel>
              </v-expansion-panels>
            </v-row>
            <div class="text-center">
              <v-pagination
                v-model="page"
                :length="Math.ceil(buyerHistory.transactions.length/perPage)"
              ></v-pagination>
            </div>

            <!-- Other buyers -->
            <br>
            <h3>Other Buyers With The Same IP</h3>
            <v-data-table
              :headers="[{text: 'ID', value: 'id'}, {text: 'Name', value: 'name'}, {text: 'Age', value: 'age'}]"
              :items="buyerHistory.otherbuyers"
              :items-per-page="5"
              class="elevation-1"
              :footer-props="{
                itemsPerPageOptions:[5,10,15,20]
              }"
            ></v-data-table>

            <!-- Recommended products -->
            <br>
            <h3>Recommended Products</h3>
            <v-simple-table>
              <template v-slot:default>
                <thead>
                  <tr>
                    <th class="text-left">
                      ID
                    </th>
                    <th class="text-left">
                      Name
                    </th>
                    <th class="text-left">
                      Price
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="(product,i) in buyerHistory.recommendedproducts"
                    :key="i"
                  >
                    <td>{{ product.id }}</td>
                    <td>{{ product.name }}</td>
                    <td>${{ product.price/100 }}</td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
<!-- End of main content -->
        </div>
      </div>
    </div>
  </v-container>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Buyerhistory',
  data () {
    return {
      id: this.$route.params.id,
      buyerHistory: {},
      loading: true,
      errored: false,
      page: 1,
      perPage: 5,
      input: null
    }
  },
  mounted () {
    this.getData()
  },
  watch: {
    $route (to, from) {
      this.id = this.$route.params.id
      this.loading = true
      this.errored = false
      this.getData()
    }
  },
  methods: {
    goToBuyer: function () {
      if (this.input !== null && this.input !== undefined && this.input !== '') {
        this.$router.push('/buyer/' + this.input)
        this.input = ''
      }
    },
    getData: function () {
      if (this.id !== undefined) {
        axios
          .get(' http://localhost:5000/api/buyer?id=' + this.id)
          .then(response => {
            this.buyerHistory = response.data
          })
          .catch(error => {
            console.log(error)
            this.errored = true
          })
          .finally(() => {
            this.loading = false
          })
      }
    }
  }

}
</script>
