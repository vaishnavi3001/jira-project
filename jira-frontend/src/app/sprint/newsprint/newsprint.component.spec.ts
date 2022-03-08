import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewsprintComponent } from './newsprint.component';

describe('NewsprintComponent', () => {
  let component: NewsprintComponent;
  let fixture: ComponentFixture<NewsprintComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewsprintComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewsprintComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
