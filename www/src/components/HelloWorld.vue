<template>
  <v-container>
    <v-layout text-xs-center wrap>
      <v-flex xs12>
        <v-img
          :src="require('../assets/logo-plamienok.png')"
          class="my-3"
          contain
          height="170"
        ></v-img>
      </v-flex>

      <v-flex>
        <h1 class="display-2 font-weight-bold mb-3">
          Plamienok App
        </h1>
        <p class="subheading font-weight-regular">
          Interná aplikácia pre neziskovú organizáciu Plamienok
          <br>
          <a href="https://www.plamienok.sk" target="_blank">www.plamienok.sk</a>
        </p>
      </v-flex>

      <v-flex xs12 mb-5>
        <h2 class="headline font-weight-bold mb-3">Štatistika</h2>
          <v-layout justify-center>
            <v-flex xs12 sm7 md5 lg4>
              <v-card>
                <v-card-title><h4>Štatistika</h4></v-card-title>
                <v-divider></v-divider>
                <v-list dense>
                  <v-list-tile v-for="stat in stats">
                    <v-list-tile-content>{{stat.text}}:</v-list-tile-content>
                    <v-list-tile-content class="align-end">{{stat.count}}</v-list-tile-content>
                  </v-list-tile>
                </v-list>
              </v-card>
            </v-flex>
          </v-layout>
      </v-flex>

      <v-flex xs12>
        <h2 class="headline font-weight-bold mb-3">Aplikácia</h2>

        <v-layout justify-center>
          Aplikácia je Open Source aplikácia. Zdrojové kódy:&nbsp;<a href="https://github.com/malaschitz/plamienok" target="_blank">
            https://github.com/malaschitz/plamienok</a>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  export default {
    data: () => ({
      stats: [
        {
          text: 'Počet užívateľov',
          count: 20
        },
        {
          text: 'Počet klientov',
          count: 10
        },
        {
          text: 'Výjazdy za minulý týždeň',
          count: 10
        },
      ]
    })
  }
</script>

<style>

</style>
