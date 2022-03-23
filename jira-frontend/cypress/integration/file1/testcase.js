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
        cy.visit('http://localhost:4200/home/1/issues')
        cy.contains('Sprint')
       

      })
      it('View Sprints', () => {
        cy.visit('http://localhost:4200/home/sprint/list')
        cy.contains('Sprint1')
        cy.contains('Create delete button')
        cy.contains('Create Add button')
        cy.contains('Sprint2')
        cy.contains('Create add API')
        cy.contains('Create Delete API')
        cy.contains('Write test cases')
       

      })
      
      
        
        
    
  })