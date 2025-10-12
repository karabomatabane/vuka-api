import { Component, computed, Input, inject } from '@angular/core';
import { signal } from '@angular/core';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { CommonModule } from '@angular/common';
import { Router, RouterLink, RouterLinkActive } from '@angular/router';
import { MatTooltipModule } from '@angular/material/tooltip';
import { AuthenticationService } from 'src/app/_services/auth.service';

export type MenuItem = {
  icon: string;
  label: string;
  route: string;
};

@Component({
  selector: 'app-x-sidenav',
  standalone: true,
  imports: [
    CommonModule,
    MatListModule,
    MatIconModule,
    RouterLink,
    RouterLinkActive,
    MatTooltipModule,
  ],
  templateUrl: './x-sidenav.component.html',
  styleUrls: ['./x-sidenav.component.scss'],
})
export class XSidenavComponent {
  private authService = inject(AuthenticationService);
  private router = inject(Router);

  currentUser = this.authService.currentUser;

  sideNavCollapsed = signal(false);
  @Input() set collapsed(val: boolean) {
    this.sideNavCollapsed.set(val);
  }

  profilePicSize = computed(() => (this.sideNavCollapsed() ? '32' : '100'));

  menuItems = computed(() => {
    if (this.currentUser()) {
      return [
        { icon: 'dashboard', label: 'Dashboard', route: '/dashboard' },
        { icon: 'newsstand', label: 'Articles', route: '/articles' },
        { icon: 'home_storage', label: 'Sources', route: '/sources' },
        { icon: 'folder', label: 'Directories', route: '/directories' },
        {
          icon: 'security',
          label: 'Roles & Permissions',
          route: '/roles-and-permissions',
        },
      ];
    } else {
      return [{ icon: 'login', label: 'Login', route: '/login' }];
    }
  });

  logout() {
    this.authService.logout();
    this.router.navigate(['/login']);
  }
}
