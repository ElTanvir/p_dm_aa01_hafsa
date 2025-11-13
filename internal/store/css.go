package store

var CriticalCSS = `
<style>
      :root {
        /* Fonts */
        /* Mapped to Design System: 'Inter' (sans) and 'Nunito' (display) */
        --font-body: "Inter", sans-serif;
        --font-title: "Nunito", sans-serif;

        /* Light Theme Colors (Warm Honey Palette) */
        --color-surface: #FFF0BF; /* Given Background */
        --color-surface-alt: #F7E9B9; /* Slightly darker, muted surface */
        --color-on-surface: #5C280B; /* Readable dark brown text */
        --color-on-surface-strong: #3E1B07; /* Stronger dark brown text */

        --color-primary: #84390F; /* Given Brand/Primary */
        --color-on-primary: #FFFFFF; /* Text on primary */

        --color-secondary: #EAAA00; /* Rich Gold accent */
        --color-on-secondary: #3E1B07; /* Dark text on gold */

        --color-outline: #DDCFA9; /* Muted beige border */
        --color-outline-strong: #84390F; /* Primary color border */

        --color-info: #0F5E84; /* Deep, muted blue */
        --color-on-info: #FFFFFF;

        --color-success: #386641; /* Earthy / Forest Green */
        --color-on-success: #FFFFFF;

        --color-warning: #D95F00; /* Burnt Orange */
        --color-on-warning: #FFFFFF;

        --color-danger: #C62828; /* Deep, earthy red */
        --color-on-danger: #FFFFFF;

        /* Border Radius */
        --radius-none: 0;
        --radius-radius: 1rem; /* Mapped to borderRadius-brand */

        /* Shadows (Tinted with Primary Color) */
        --shadow-sm: 0 1px 2px 0 rgba(132, 57, 15, 0.05);
        --shadow-md: 0 4px 10px rgba(132, 57, 15, 0.07); /* Mapped to boxShadow-brand */
        --shadow-lg: 0 10px 15px -3px rgba(132, 57, 15, 0.1),
          0 4px 6px -4px rgba(132, 57, 15, 0.07);
        --shadow-xl: 0 20px 25px -5px rgba(132, 57, 15, 0.1),
          0 8px 10px -6px rgba(132, 57, 15, 0.1);
      }
    </style>
  `
