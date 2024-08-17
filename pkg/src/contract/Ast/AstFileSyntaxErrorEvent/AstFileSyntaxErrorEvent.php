<?php

declare (strict_types=1);
namespace Qossmic\Deptrac\Contract\Ast;

use DEPTRAC_INTERNAL\Symfony\Contracts\EventDispatcher\Event;
/**
 *
 */
final class AstFileSyntaxErrorEvent extends Event
{
    public function __construct(public readonly string $file, public readonly string $syntaxError)
    {
    }
}
