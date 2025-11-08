import { Component, input, signal, OnDestroy } from '@angular/core';
import { MatTooltipModule } from '@angular/material/tooltip';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { MenuItem } from '../x-sidenav/x-sidenav.component';
import { MatIcon } from '@angular/material/icon';
import { CommonModule } from '@angular/common';
import { MatListModule } from '@angular/material/list';
import { MatButtonModule } from '@angular/material/button';
import { animate, style, transition, trigger } from '@angular/animations';

@Component({
  selector: 'app-menu-item',
  standalone: true,
  animations: [
    trigger('expandCollapseMenu', [
      transition(':enter', [
        style({ opacity: 0, height: '0px' }),
        animate('500ms ease-in-out', style({ opacity: 1, height: '*' })),
      ]),
       transition(':leave', [
        animate('500ms ease-in-out', style({ opacity: 0, height: '0px' })),
      ]),
    ]),
  ],
  imports: [
    CommonModule,
    RouterLink,
    RouterLinkActive,
    MatTooltipModule,
    MatListModule,
    MatIcon,
    MatButtonModule,
  ],
  templateUrl: './menu-item.component.html',
  styleUrl: './menu-item.component.scss',
})
export class MenuItemComponent implements OnDestroy {
  item = input.required<MenuItem>();
  collapsed = input<boolean>(false);
  onClick = input<() => void>();
  visible = input<boolean>(true);
  nestedMenuOpen = signal<boolean>(false);
  private menuHoverTimeout?: ReturnType<typeof setTimeout>;
  private isMenuClicked = false;

  toggleNestedMenu() {
    if (!this.item().subItems) {
      return;
    }
    this.nestedMenuOpen.set(!this.nestedMenuOpen());
  }

  setNestedMenuState(isOpen: boolean) {
    if (!this.item().subItems) {
      return;
    }
    
    // Clear any existing timeout
    if (this.menuHoverTimeout) {
      clearTimeout(this.menuHoverTimeout);
      this.menuHoverTimeout = undefined;
    }

    if (isOpen) {
      this.nestedMenuOpen.set(true);
    } else {
      // Don't close immediately if menu was clicked
      if (this.isMenuClicked) {
        this.isMenuClicked = false;
        return;
      }
      
      // Add a small delay before closing to prevent accidental closes
      this.menuHoverTimeout = setTimeout(() => {
        this.nestedMenuOpen.set(false);
      }, 150);
    }
  }

  onMenuItemClick() {
    if (this.item().subItems) {
      this.isMenuClicked = true;
      this.setNestedMenuState(true);
    }
  }

  onSubItemClick() {
    this.isMenuClicked = true;
  }

  ngOnDestroy() {
    if (this.menuHoverTimeout) {
      clearTimeout(this.menuHoverTimeout);
    }
  }
}
