/// <reference types="cypress" />

describe('My First Test', () => {
    it('Does not do much!', () => {
      expect(true).to.equal(true)
    })
    it('Visit homepage', () => {
        cy.visit('http://localhost:4200/')
        cy.contains('Create')
        cy.contains('Your Work')
        cy.contains('People')
        cy.contains('Filters')
        cy.contains('Projects')
        cy.contains('App')

      })
      it('View issues', () => {
        cy.visit('http://localhost:4200/home/project/1/issues')
        cy.contains('Sprint')
       

      })
      it('View Sprints', () => {
        cy.visit('http://localhost:4200/home/project/1/sprints')
        cy.contains('Sprint 1')
        // cy.contains('Create delete button')
        // cy.contains('Create Add button')
        cy.contains('Sprint 2')
        cy.contains('Sprint 4')
        // cy.contains('Create add API')
        // cy.contains('Create Delete API')
        // cy.contains('Write test cases')
       

      })
      
      it('Login', () => {
        cy.visit('http://localhost:4200/login')
        cy.contains('Username')
        cy.contains('Password')
        cy.contains('Login')
     
        
       

      })

      it('Register', () => {
        cy.visit('http://localhost:4200/register')
        cy.contains('SIGN UP')
        cy.contains('First Name')
        cy.contains('Last Name')
        cy.contains('Email')
        cy.contains('Password')
        cy.contains('Username')
        
       

      })
      

      
      
      
        
        
    
  })