import { constant_values } from "../../constants";

describe("Checks the issue-detail view", () => {

    it("should test if he issue detail view opens", () => {
    
      localStorage.setItem('access-token', `${constant_values.ACCESS_TOKEN}`)

      cy.visit("http://localhost:4200/home/issue/1/details");
      cy.wait(2000);
    
    });
  
  });