package cookiemonster

import "log"

const (
	AsciiArt = ` ::::::::   ::::::::   ::::::::  :::    ::: ::::::::::: ::::::::::              
:+:    :+: :+:    :+: :+:    :+: :+:   :+:      :+:     :+:                     
+:+        +:+    +:+ +:+    +:+ +:+  +:+       +:+     +:+                     
+#+        +#+    +:+ +#+    +:+ +#++:++        +#+     +#++:++#                
+#+        +#+    +#+ +#+    +#+ +#+  +#+       +#+     +#+                     
#+#    #+# #+#    #+# #+#    #+# #+#   #+#      #+#     #+#                     
 ########   ########   ########  ###    ### ########### ##########              
::::    ::::   ::::::::  ::::    :::  :::::::: ::::::::::: :::::::::: ::::::::: 
+:+:+: :+:+:+ :+:    :+: :+:+:   :+: :+:    :+:    :+:     :+:        :+:    :+:
+:+ +:+:+ +:+ +:+    +:+ :+:+:+  +:+ +:+           +:+     +:+        +:+    +:+
+#+  +:+  +#+ +#+    +:+ +#+ +:+ +#+ +#++:++#++    +#+     +#++:++#   +#++:++#: 
+#+       +#+ +#+    +#+ +#+  +#+#+#        +#+    +#+     +#+        +#+    +#+
#+#       #+# #+#    #+# #+#   #+#+# #+#    #+#    #+#     #+#        #+#    #+#
###       ###  ########  ###    ####  ########     ###     ########## ###    ###`
)

func LogAsText(cookies []Cookie) error {
	for _, cookie := range cookies {
		log.Printf("[+] Host: %s\n", cookie.Host)
		log.Printf("    Name: %s\n", cookie.Name)
		log.Printf("    Value: %s\n", cookie.Value)
		log.Printf("    Path: %s\n", cookie.Path)
		log.Printf("    IsSecure: %t\n", cookie.IsSecure)
		log.Printf("    IsHttpOnly: %t\n", cookie.IsHttpOnly)
		log.Printf("    CreationUtc: %d\n", cookie.CreationUtc)
		log.Printf("    ExpiryUtc: %d\n", cookie.ExpiryUtc)
		log.Println("----------------------------------------------------")
	}

	return nil
}
