import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TrialDragAndDropComponent } from './trial-drag-and-drop.component';

describe('TrialDragAndDropComponent', () => {
  let component: TrialDragAndDropComponent;
  let fixture: ComponentFixture<TrialDragAndDropComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TrialDragAndDropComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TrialDragAndDropComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
