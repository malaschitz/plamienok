<template>
  <v-container class="gray lighten-3">
    <v-row justify="center" align="center">
      <v-col cols="12" sm="12" align="center">
        <v-img :src="require('../assets/logo-plamienok.png')" class="my-1" contain width="300" />
      </v-col>
    </v-row>

    <v-row v-if="!$store.state.logged">
      <v-col align="center">
        <h2 class="display-1">LOGIN</h2>
        <v-card min-width="400" max-width="600" v-if="!forgotMode" class="pa-3">
          <v-alert v-model="showAlert" type="error" prominent :dismissible="true">{{ alertMessage }}</v-alert>
          <v-form v-model="valid" lazy-validation ref="form" class="ma-2">
            <v-text-field label="Email" v-model="email" :rules="emailRules" required />
            <v-text-field
              label="Heslo"
              v-model="password"
              required
              :append-icon="show1 ? 'mdi-eye-off' : 'mdi-eye'"
              :type="show1 ? 'text' : 'password'"
              name="input-10-1"
              counter
              @click:append="show1 = !show1"
            />
            <v-btn color="success" @click="login" class="ma-3">Prihlásiť</v-btn>
            <v-btn color="info" @click="forgot" class="ma-3">Zabudnuté heslo</v-btn>
          </v-form>
        </v-card>

        <v-card min-width="400" max-width="600" v-if="forgotMode">
          <v-alert v-model="showAlert" type="error" prominent :dismissible="true">{{ alertMessage }}</v-alert>
          <v-form v-model="valid" lazy-validation ref="formCode" class="ma-2">
            <v-text-field
              label="Kód z emailu"
              v-model="code6"
              :rules="code6Rules"
              required
              mask="######"
            />

            <h2>Nové heslo (nepovinné)</h2>

            <v-text-field
              label="Heslo"
              v-model="password"
              :append-icon="show2 ? 'visibility_off' : 'visibility'"
              :type="show2 ? 'text' : 'password'"
              name="input-10-1"
              counter
              @click:append="show2 = !show2"
            />

            <v-text-field
              label="Heslo"
              v-model="password2"
              :append-icon="show3 ? 'visibility_off' : 'visibility'"
              :type="show3 ? 'text' : 'password'"
              name="input-10-1"
              counter
              @click:append="show3 = !show3"
            />

            <v-btn color="success" @click="sendCode6">Poslať</v-btn>
            <v-btn color="info" @click="back">Späť</v-btn>
          </v-form>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col align="center">
        <h3 class="display-1">Plamienok App</h3>
        <p class="subheading font-weight-regular">
          Interná aplikácia pre neziskovú organizáciu Plamienok
          <br />
          <a href="https://www.plamienok.sk" target="_blank">www.plamienok.sk</a>
        </p>
      </v-col>
    </v-row>

    <v-row>
      <v-col align="center">
        <h3 class="display-1">Štatistika</h3>
        <v-card min-width="300" max-width="400">
          <v-card-title>
            <h4>Štatistika</h4>
          </v-card-title>
          <v-divider />
          <v-list dense>
            <v-list-item v-for="stat in stats" :key="stat.Text">
              <v-list-item-content align="left">
                <v-list-item-title v-text="stat.Text"></v-list-item-title>
              </v-list-item-content>
              <v-list-item-content align="right">
                <v-list-item-title v-text="stat.Value"></v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col align="center" class="gray lighten-3">
        <h3 class="display-1">Aplikácia</h3>

        <p>
          Aplikácia je Open Source aplikácia.
          <br />
          <v-btn to="/about">Viac o aplikácii</v-btn>
        </p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Dashboard",

  data: () => ({
    valid: false,
    email: "",
    emailRules: [
      v => !!v || "Email je vyžadovaný",
      v => /.+@.+/.test(v) || "Email musí mať správny formát"
    ],
    password: "",
    password2: "",
    stats: [],
    forgotMode: false,
    show1: false,
    show2: false,
    show3: false,
    alertMessage: "",
    showAlert: false,
    code6: "",
    code6Rules: [
      v => !!v || "Kód musí byť zadaný",
      v => /^[0-9]{6}$/.test(v) || "Kód je 6 číselných znakov"
    ]
  }),

  methods: {
    login: function() {
      this.showAlert = false;
      this.alertMessage = "";

      if (this.valid && this.email !== "" && this.password !== "") {
        this.$axios
          .post("/api/login", {
            email: this.email,
            password: this.password
          })
          .then(response => {
            console.log("OK", response);
            if (response.status == 202) {
              this.showAlert = true;
              this.alertMessage = "Chyba: " + response.data.Error;
            } else {
              this.loginSave(response.data);
            }
          })
          .catch(response => {
            console.log("WRONG", response);
            this.showAlert = true;
            this.alertMessage = "Chyba: " + response;
          });
      } else {
        this.showAlert = true;
        this.alertMessage = "Chýbajú údaje na prihlásenie";
      }
    },

    forgot: function() {
      this.showAlert = false;
      this.alertMessage = "";
      if (this.email !== "") {
        console.log("Forgot", this.email);
        this.$axios
          .post("/api/forgot", { email: this.email })
          .then(response => {
            if (response.status == 202) {
              this.showAlert = true;
              this.alertMessage = "Chyba: " + response.data.Error;
            } else {
              this.password = "";
              this.password2 = "";
              this.code6 = "";
              this.forgotMode = true;
            }
          })
          .catch(response => {
            console.log("WORNG", response);
            this.showAlert = true;
            this.alertMessage = "Chyba: " + response;
          });
      } else {
        this.alertMessage = "Email musí byť zadaný";
        this.showAlert = true;
      }
    },

    sendCode6: function() {
      this.showAlert = false;
      this.alertMessage = "";
      if (this.code6 === "") {
        this.alertMessage = "Kód musí byť zadaný";
        this.showAlert = true;
      } else if (this.password !== this.password2) {
        this.alertMessage = "Opakované heslo nie je rovnaké";
        this.showAlert = true;
      } else if (this.password.length > 0 && this.password.length < 6) {
        this.alertMessage = "Heslo je príliš krátke";
        this.showAlert = true;
      } else {
        this.$axios
          .post("/api/checkCode6", {
            email: this.email,
            code6: this.code6,
            password: this.password
          })
          .then(response => {
            if (response.status == 202) {
              this.showAlert = true;
              this.alertMessage = "Chyba: " + response.data.Error;
            } else {
              if (response.status == 202) {
                this.showAlert = true;
                this.alertMessage = "Chyba: " + response.data.Error;
                this.forgotMode = true;
              } else {
                this.loginSave(response.data);
              }
            }
          })
          .catch(response => {
            console.log("WORNG", response);
            this.showAlert = true;
            this.alertMessage = "Chyba: " + response;
          });
      }
    },

    loginSave: function(data) {
      console.log("loginSave", data.User);
      this.$store.commit("login", {
        name: data.User.Name,
        id: data.User.ID,
        email: data.User.Email
      });
      localStorage.token = data.Token;
      this.$axios.defaults.headers.common["token"] = localStorage.token;
      this.forgotMode = false;
    },

    back: function() {
      this.forgotMode = false;
    }
  },

  mounted: function() {
    console.log("mounted 2");
    this.$axios
      .get("/api/stats")
      .then(response => {
        console.log(response);
        this.stats = response.data;
      })
      .catch(response => {
        console.log("Chyba", response);
      });
  }
};
</script>

<style scoped></style>
