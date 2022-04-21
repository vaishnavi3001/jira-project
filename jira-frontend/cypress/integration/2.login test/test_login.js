import { constant_values } from "../../constants";
// import 'cypress-localstorage-commands'


describe("Checks Login", () => {

    it("should test the login page", () => {
      cy.visit(`${constant_values.LOGIN}`);
      cy.get("[id=email]").type("pypalkar23");
      cy.get("[id=password]").type("mandar");
  
      cy.get("[id=login-button]").click();
      cy.wait(2000);

    //   cy.getCookie('access-token')
    // .then((cookie) => {
    //   cy.log(cookie)
    //   cy.setCookie("access-token", cookie.value, { path: '/' })
    // })
      // cy.saveLocalStorage();
      //localStorage.setItem('access_token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NTA1MDA3OTQsImlhdCI6MTY1MDQ5NzE5NH0.c14W1sS-vjg9lYWY03GPKsnjkGzPPJrrb_6KTgkXlL0')
    });
  
  });