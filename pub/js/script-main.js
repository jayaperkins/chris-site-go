$(function() {
    
    var $body = $('body');
    
    $body.hide(0).delay(500).fadeIn(2500);
    
    var $linkContainer = $('.link-container');
    
    $linkContainer.on('mouseover', function() {
        
        $(this).addClass('hover');
        
    });
    
    $linkContainer.on('mouseout', function() {
        
        $(this).removeClass('hover');
        
    });
});